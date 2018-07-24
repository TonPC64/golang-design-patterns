package adapter

import "testing"

func TestAdapter(t *testing.T) {
	msg := "HelloWorld!"
	adapter := PrinterAdapter{
		OldPrinter: &MyLegacyPrinter{},
		Msg:        msg,
	}

	returnedMsg := adapter.PrintStored()

	if returnedMsg != "Legacy Printer: Adapter: HelloWorld!\n" {
		t.Errorf("Message didn't match: %s\n", returnedMsg)
	}

	adapter = PrinterAdapter{
		OldPrinter: nil,
		Msg:        msg,
	}

	returnedMsg = adapter.PrintStored()

	if returnedMsg != "HelloWorld!" {
		t.Errorf("Message didn't match: %s\n", returnedMsg)
	}

}
