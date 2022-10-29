### Domain ###
- User
```
{
    "id": B,
    "fullName": "User B",
    "userName": "markhaur",
    "groups": [G1,G2],
    "contacts": [A,C]
}
```
- Group
```
{
    "id": G2,
    "name": ...,
    "participants": [B, C]
}
```
- MessageSent
```
{
    "id": "123"
    "sender": ...,
    "content": ...,
    "createdAt": ...,
    "sentAt": ...,
    "recipients": {
        "B": [
            {"in": "G1", "delivered": "12:00:00", "seen": "12:01:05"},
            {"in": "G2", "delivered": "12:00:00", "seen": "12:01:09"}
        ],
        "C": [
            {"in": "G2", "delivered": "12:00:00", "seen": "12:01:05"},
            {"in": "DM", "delivered": "12:00:00", "seen": "12:01:09"}
        ],
        "D": [
            {"in": "DM", "delivered": "12:00:00", "seen": "12:01:09"}
        ]
    }
}
```
- MessageReceived
```
// By B:
{
    "id": "123",
    "in": [G1,G2],
    "messageFrom": A,
    "content": ..., 
    "createdAt": ...,
    "sentAt": ...,
    "receivedAt": ...
}
```
```
// By C:
{
    "id": "123",
    "in": [DM*,G2],
    "messageFrom": A,
    "content": ..., 
    "createdAt": ...,
    "sentAt": ...,
    "receivedAt": ...
}
// *DM stands for Direct Message
```
```
// By D:
{
    "id": "123",
    "in": [DM],
    "messageFrom": A,
    "content": ..., 
    "createdAt": ...,
    "sentAt": ...,
    "receivedAt": ...
}
```
### Usecases ###
- User A can send a Message to User B
- User B can send a Message to User A
- User A can send a Message to User B and User C (like forwarding a link to 2 users)
- User A can send a Message to a Group (User B and User C)

MongoDB 3 instances
if messageFrom == A
    should we save locally?
    or save remotely as well?

instancesByUser := map[string]int{
    "A": 2,
    "B": 1,
    "C": 3,
}

C->A
save in (instancesByUser(A)) // 2
save in (instancesByUser(C)) // 3

A->B
save in (instancesByUser(B)) // 1
save in (instancesByUser(A)) // 2

Duplication?



CS1 CS2
^    ^
A -> B

CS1 -> CS2
event



Domain <- Application <- Infrastructure

Sender----> CS1                     Consumer                    CS2                                           <----Recipient
Send        Send, PublishSentEvent  HandleSentEvent, Relay      Relay, PublishRelayEvent, HandleRelayEvent
Websocket   Websocket, NATS         NATS             Websocket  Websocket, Go channel     Websocket                Websocket