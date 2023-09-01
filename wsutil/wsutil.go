package wsutil

import "github.com/thanos-go/log"

func LogSend(accountId uint, event string, additions map[string]interface{}) {
	if additions == nil {
		additions = make(map[string]interface{})
	}
	additions["event"] = event
	logTransmission(accountId, true, additions)
}

func LogReceive(accountId uint, action string, additions map[string]interface{}) {
	if additions == nil {
		additions = make(map[string]interface{})
	}
	additions["action"] = action
	logTransmission(accountId, false, additions)
}

func logTransmission(accountId uint, isSend bool, additions map[string]interface{}) {

	fields := map[string]interface{}{
		"account_id": accountId,
		"status":     "ok",
		"kind":       "receive",
	}

	for key, field := range additions {
		if key == "error" {
			fields["status"] = "error"
		}
		fields[key] = field
	}

	if isSend {
		fields["kind"] = "send"
	}

	log.Info("websocket activity", fields)
}
