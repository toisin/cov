package workflow

type CovPrompt struct {
	previousPrompt Prompt
	// response Response
	expectedResponseHandler *ExpectedResponseHandler
	// promptGenerator PromptGenerator
	currentUIPrompt UIPrompt
	promptConfig PromptConfig
}

func MakeCovPrompt(p PromptConfig) *CovPrompt {
	// pg := new(CovPromptGenerator)
	// pg.promptID = p.Id
	// pg.uiActionModeId = p.UIActionModeId
	// pg.promptText = p.Text

	erh := MakeExpectedResponseHandler(p.ExpectedResponses)

	n := new(CovPrompt)
	n.promptConfig = p
	n.expectedResponseHandler = erh
	return n
}

// func (cp *CovPrompt) GetDisplayText() string {
// 	return cp.promptGenerator.GetPromptText()
// }

// func (cp *CovPrompt) GetUIActionModeId() string {
// 	return cp.promptGenerator.GetUIActionModeId()
// }

func (cp *CovPrompt) GetPhaseId() string {
	return PHASE_COV
}

// func (cp *CovPrompt) SetResponse(r Response) {
// 	cp.response = r
// }

// func (cp *CovPrompt) GetNextPrompt() Prompt {
// 	return cp.expectedResponseHandler.GetNextPrompt(cp.response)
// }

// func (cp *CovPrompt) GetResponseText() string {
// 	return cp.response.GetText()
// }

func (cp *CovPrompt) GetUIPrompt() UIPrompt {
	if (cp.currentUIPrompt == nil) {
		pc := cp.promptConfig
		switch pc.Type {
		case UI_PROMPT_MC:
			p := NewUIMCPrompt()
			p.Text = cp.promptConfig.Text // TODO need to process dynamic data
			// p.WorkflowStateID = "1" // TODO clean up legacy
			p.Options = make([]UIOption, len(pc.ExpectedResponses))
			for i := range pc.ExpectedResponses {
				p.Options[i] = UIOption{pc.ExpectedResponses[i].Text, pc.ExpectedResponses[i].Id}
			}
			p.PromptId = pc.Id
			cp.currentUIPrompt = p
			break
		case UI_PROMPT_TEXT:
			p := NewUITextPrompt()
			p.Text = cp.promptConfig.Text // TODO need to process dynamic data
			// p.WorkflowStateID = "1" // TODO clean up legacy
			p.PromptId = pc.Id
			cp.currentUIPrompt = p
			break
		}
	}
	return cp.currentUIPrompt
}


// type CovPromptGenerator struct {
// 	promptID string
// 	uiActionModeId string // the mode of rendering for Action UI
// 	promptText string // text with place holders for dynamic data
// 	state State
// 	previousPrompt Prompt
// }

// func (cph *CovPromptGenerator) GenerateUIPrompt() UIPrompt {
// 	//TODO hardcoding the prompt
// 	return &UIMCPrompt{
// 						UI_PROMPT_MC,
// 						"1",
// 						"Let's get started! What feature have you investigated?",
// 						"",
// 					 	[]UIOption{
// 					 		UIOption{variableMap["X1"].Text,"X1"},
// 					 		UIOption{variableMap["X2"].Text,"X2"},
// 					 		UIOption{variableMap["X3"].Text,"X3"},
// 					 		UIOption{variableMap["X4"].Text,"X4"}},
// 					 	"p1"}
// 	// return &UITextPrompt{UI_PROMPT_YES_NO, "2", "Do you think it makes a difference?", "1", "3"}
// 	// return &UITextPrompt{
// 	// 					UI_PROMPT_TEXT,
// 	// 					"3",
// 	// 					cp.GetDisplayText(),
// 	// 					"2",
// 	// 					"4"}
// 	// return = &UITextPrompt{UI_PROMPT_TEXT, "4", "What did you find out about %X1?", "3", "5"}
// 	// return = &UITextPrompt{UI_PROMPT_TEXT, "5", "How do you know?", "4", "6"}
// 	// return = &UITextPrompt{UI_PROMPT_TEXT, "6", "Which records show you are right?", "5", UI_PROMPT_END}
// 	// return = &UITextPrompt{UI_PROMPT_END, UI_PROMPT_END, "You have done!", "6", UI_PROMPT_END}
// }

// func (cph *CovPromptGenerator) generatePromptText() string {
// 	// TODO
// 	// data = state.GetDynamicData() // Get needed dynamic data from state
// 	// text = generatePromptText(data)
// 	cph.promptText = ""
// 	return cph.promptText
// }

// func (cph *CovPromptGenerator) generateAction() string {
// 	// TODO
// 	// data = state.GetDynamicData() // Get needed dynamic data from state
// 	// text = generatePromptText(data)
// 	return cph.uiActionModeId
// }

// func (cph *CovPromptGenerator) GetPromptText() string {
// 	if (cph.promptText == "") {
// 		cph.generatePromptText()
// 	}
// 	return cph.promptText
// }

// func (cph *CovPromptGenerator) GetUIActionModeId() string {
// 	return cph.uiActionModeId
// }
