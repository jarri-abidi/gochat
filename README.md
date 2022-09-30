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
    "toGroups": [G1]
    "toUsers": [B]
	"content": ... // imp for B
	"createdAt": ...
    "sentAt": ...
    "received": [{"by": B, "at": "12:00:00"}] // imp for A
    "seen": [{"by": B, "at": "12:00:00"}] // imp for A
}
```
- MessageReceived
```
{
    "messageId": 123
    // "chatId": ...
	"messageFrom": A
	"content": ... // imp for B
	"createdAt": ...
    "sentAt": ...
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
