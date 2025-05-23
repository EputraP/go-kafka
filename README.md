# go-kafka
- kafka consist of topic, producer, and consumer
- the producer and consumer should communicate in the same topic
- topic will stored the message that sent by producer
- The topic structure is similar to a log, where new messages are appended to the end.
- If the consumer doesn't define the group ID, Kafka will assign one randomly.
- message structure:
    - topic     : The name of the Kafka topic to which the message is sent.
    - partition : The specific partition within the topic where the message is stored.
    - header    : Optional metadata in the form of key-value pairs. Useful for things like tracing IDs, content types, etc.
    - key       : A key used to determine the target partition. It's optional and is in byte format.
    - value     : The main content or payload of the message, also in byte format. It’s typically encoded (e.g., JSON, Avro, Protobuf).
    - offset    : A unique ID within the partition that identifies the position of the message.
    - timestamp : The time the message was produced or written to Kafka (depending on configuration).
- the key is optional but,:
    - not defined
        - it uses a round-robin strategy to distribute messages evenly across all partitions.
        - No Ordering Guarantee (Kafka only guarantees ordering within a single partition.)
        - Better load balancing, but no order
    - defined same value
        - Messages go to same partition and are stored as new records
        - Order is preserved per key
    - defined different value
        - Likely different partitions (Partition chosen by key hash)
        - ordering is guaranteed within each partition, and therefore per key
            - ex 
            ```json 
            [
                {
                    key:key1,
                    value:123
                },{
                    key:key1,
                    value:1234
                },{
                    key:key2,
                    value:12366
                },
            ]
            ```
            - the ordered will happended only on the key1
- Consumer Group
    - the consumer should defined it's consumer group
    - kafka will create automaticly if can't find the consumer group
    - usually the consumer group name is using the name of app
    - when reading the topic, kafka will only give the data to the consumer which has unique consumer group
      (it means that kafka will not give twice to the consumer that has same consumer group)
    - if the consumer group always changing, so automatically the data will be accepted in many times
    - Ex when consumer group consume from a topic with different partition:
        - if i have a topic, we called "my-topic" with 5 partition
        - and than my producer send a message to "my-topic" with different keys, so the message will distibuted to it's all partitions
        - and than i have 3 consumer
        - 2 consumer has same consumer group (consumer 1 and 2), and the other has it's own consumer group name (consumer 3)
        - the consumer 1 and 2 will shared the partition, i meant like the consumer 1 will consume partition from 1, 2, 3 and consumer 2 will consume 4, 5 maybe
        - the consumer 3 will consume all the partition, 1 until 5
    - Ex when the consumer group consume from a topic with same partition :
        - if i have a topic, we called "my-topic" with 5 partition
        - and than my producer send a message to "my-topic" with same keys, so the message will distibuted single partition, maybe we used partition 2
        - and than i have 3 consumer
        - 2 consumer has same consumer group (consumer 1 and 2), and the other has it's own consumer group name (consumer 3)
        - if the consumer 2 already consume the message, so the consumer 1 will not got the message
        - the consumer 3 will consume all the partition, 1 until 5
    - if the consumer in the same consumer group, one of them offline, the other will back up
        - ex:
        - there are 3 consumers, and if consume in the same partition actually the online ones will consume the message
        -  if the different partition, the message will be distibuted only for the online consumer 
    - additional information for different parition
        -  if example there are 2 consumer with same consumer group
        - if the consumer 1 already consumed all the messages, so the consumer 2 will got nothing
        - but if not, the message will be shared, but when i tried, it looks like not in the same amount 
    - additional information for same partition:
        - it suitables for when there are more then one service which has same function
- offset
    - An offset is simply a increasing number assigned to each message within a partition.
    - It represents the position of a message within the partition.
    - Offset starts at 0 for each partition.
    - Every time a message is appended to the partition, Kafka increases the offset by 1.
    - Each partition has its own independent offset counter
    - Offsets are never reused — once assigned, they’re fixed
    - They are assigned at the moment the message is written
    - Offset Commit
        - When a Kafka consumer reads a message, it doesn't mean the message is “done” — it only means it’s been delivered to the consumer.
            - Processed it successfully
            - Crashed while processing
            - Skipped it
        - To solve this, Kafka lets consumers "commit the offset", which means: "I’ve successfully processed all messages up to this offset, you can safely move on."
        - Kafka stores committed offsets in an internal topic: "__consumer_offsets"
        - simply offset commit is tell the kafka the consumer last consumed offset as its checkpoint
        - basically, we can set the start of the offset menually
            - there are 2 types
            - with consumer group or not
            - with consumer group
                - if using consumer group (group id), you only be able to set for the s kafka.firstOffset, lastOffset, and FirstOffset
            - without consumer group
                - you should define the partition
                - you can assign an integer value to StartOffset belongs to the offset that you want
                - you also can use the SetOffset
        - there are 2 types of offset commit (manual and auto):
            - auto-commit
                - Kafka will automatically commit the latest offset at a set interval
                - The default is every 5 seconds (auto.commit.interval.ms = 5000).
                - If your app crashes after reading but before processing, Kafka will assume the message was handled and will not deliver it again = ❗potential data loss.
            - manual commit
                - You disable auto-commit and commit only after success.
                - More reliable — offset is only advanced after success
                - If the app crashes before CommitMessages(), Kafka will re-deliver the message
- Core Concept: Kafka Partitions + Consumer Groups
    - A topic is divided into partitions (units of parallelism)
    - A consumer group shares the work of consuming messages across its active consumers
    - Kafka assigns exactly one consumer per partition within a group
    - ✅ 1. Consumers = Partitions
        - each consumer gets 1 partition
        - Full parallelism, evenly distributed.
        ``` 
        3 partitions ↔ 3 consumers
        P0 → C1
        P1 → C2
        P2 → C3
        ```
    - ⚠️ 2. Consumers < Partitions
        - Some consumers get multiple partitions
        - Kafka balances as evenly as possible
        - some consumers do more work
        ``` 
        5 partitions ↔ 2 consumers
        P0, P1, P2 → C1
        P3, P4      → C2
        ```
    - ❌ 3. Consumers > Partitions
        - Some consumers are idle (get 0 partitions)
        - Kafka does not assign the same partition to multiple consumers in the same group
        ``` 
        2 partitions ↔ 4 consumers
        P0 → C1
        P1 → C2
        C3, C4 do nothing
        ```
    - 🔄 4. If a Consumer Crashes or Goes Offline
        - Kafka rebalances remaining consumers
        - Partitions of the offline consumer are reassigned
        ``` 
        Before:
        P0 → C1
        P1 → C2
        P2 → C3

        After C2 offline:
        P0 → C1
        P1 → C3
        P2 → C3
        ```
- Replicas
    - Replication means that each Kafka partition is copied to multiple brokers. These copies are called replicas
    - key conceps
        - leader    = the main broker that handles all reads/writes for a partition
        - follower  = other brokers that keep a copy (replica) for the partition
        - ISR (In-Sync Replicas)    = that set of replicas that are up to date with the leader
        - Replication Factor    =  How many total copise (leader + followers) a partition has
    - example
        ```
        Partition 0:
        - Leader → Broker 1
        - Followers → Broker 2, Broker 3

        Partition 1:
        - Leader → Broker 2
        - Followers → Broker 1, Broker 3

        Partition 2:
        - Leader → Broker 3
        - Followers → Broker 1, Broker 2
        ```
    - how replication works
        - producer writes to the leader of the partition
        - leader replicates data to followers (asyn or sync)
        - when followers catch up, they're added to the ISR
        - kafka can be configured to:
            - wait for 1 ACk (fast, but risky)
            - wait for all ISR ACKs (safer, but slower)
    - when leader goes down: kafka automatically elects a new leader from ISR
    - followers lags or dies: its temporarily removed from ISR
    - No ISR available: topic becomes unavaliable until ISR is restored
    - configuration (set on topic/producer level):
        - repilcation.factor : when creating topic
        - min.insync.replicas : minimum replicas required to acknowledge
        - producer acks setting:
            - "0" : fire and forget
            - "1" : wait for leader only
            - "all": wait for all ISR replicas
            