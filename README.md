# common TYPES AND FUNCTIONS
## contains protobuffers for grpc and as DTOs ([see here](https://github.com/aybjax/nis_kub))

## /cmntypes
  - cache - contains lower level adapter for cache. Could be redis, nats, hashmap or anything else
  - queue - lower level adapter for queue. Could be redis, nats, rabbitmq, kafka, queue data type
  - /cmntypes/mock_cmntypes - mock object by gomock; aut generated
## /helper
  - SetDiff[T comparable](first []T, second []T) (result []T)
    - Allows removal of elements of second element from first.
    - runtime.Gosched() allows context switch if elements are too big
  - GenerateUpdateMessage(modelId string, toUpdate []string, _type pbdto.UpdateType) <-chan *pbdto.UpdateEmbedded
    - Maps modelId and updated ids for 
    - All elements must be consumed for channel to be closed
    - Mimics generators and allow context switch
  - mapError implements error interface
    - used for return type of services
## pbdto and data.proto
  - Contain protobuffers, services and messages
  - Used in Grpc and as data transfer object.
