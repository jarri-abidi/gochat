### Domain ###
- User
```
{
    "id": "uuid/phone"
    "username": "markhaur",
    "groups": []
}
```
- Group
```
{
    "id": G1,
    "name": ...,
    "participants": []
}
```
- MessageSent
```
{
    "id": 123
    "toGroups": [G1, G2]
    "toUsers": [C, D]
    "content": ...,
    "createdAt": ...,
    "sentAt": ...,
    "received": [
        {"by": B, "at": "12:00:00", "in": "G1"},
        {"by": B, "at": "12:00:00", "in": "G2"},
        {"by": C, "at": "12:05:00", "in": "G2"},
        {"by": C, "at": "12:05:00"},
        {"by": D, "at": "12:00:00"}
    ] 
    "seen": [
        {"by": B, "at": "12:01:05", "in": "G1"},
        {"by": B, "at": "12:01:00", "in": "G2"},
        {"by": C, "at": "12:10:00", "in": "G2"},
        {"by": C, "at": "12:10:05"},
        {"by": D, "at": "12:20:00"}
    ] 
}
```
- MessageReceived
```
// By B:
{
    "messageId": 123,
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
    "messageId": 123,
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
    "messageId": 123,
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
