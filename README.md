# Restful-go

| HTTP verb |                 Path                 | Action | Resource |
| :-------- | :----------------------------------: | -----: | -------- |
| POST      |   /v1/train(details as JSON body)    | Create | Train    |
| POST      |  /v1/station(details as JSON body)   | Create | Station  |
| GET       |             /v1/train/id             |   Read | Train    |
| GET       |            /v1/station/id            |   Read | Station  |
| POST      | /v1/schedule(source and destination) | Create | Route    |
