# Webhook setup guide

it just plain api endpoint, 

```go
func TriggerHookExample(c fiber.Ctx) error {
    res := common.ResponseDTO[string]{
        Data:   "Hook triggered successfully",
        Status: common.SUCCESS,
        Error:  nil,
    }
    return c.Status(fiber.StatusOK).JSON(res)
}
```
