package main

/*
type LobbyMatch struct{}


type LobbyMatchState struct {
	presences  map[string]runtime.Presence
	emptyTicks int
}


func RegisterMatch(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule) (runtime.Match, error) {
	return &LobbyMatch{}, nil
}

func MatchmakerAdd(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *rtapi.Envelope) (*rtapi.Envelope, error) {
	message, ok := in.Message.(*rtapi.Envelope_MatchmakerAdd)
	if !ok {
		return nil, errInternal
	}

	// Force min count to be 4 and max count to be 8
	message.MatchmakerAdd.MinCount = 2
	message.MatchmakerAdd.MaxCount = 2

	return in, nil
}

func (m *LobbyMatch) MatchJoinAttempt(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, dispatcher runtime.MatchDispatcher, tick int64, state interface{}, presence runtime.Presence, metadata map[string]string) (interface{}, bool, string) {
	result := true

	// Custom code to process match join attempt.
	return state, result, ""
}

func (m *LobbyMatch) MatchInit(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, params map[string]interface{}) (interface{}, int, string) {
	state := &LobbyMatchState{
		emptyTicks: 0,
		presences:  map[string]runtime.Presence{},
	}
	tickRate := 1 // 1 tick per second = 1 MatchLoop func invocations per second
	label := ""
	return state, tickRate, label
}

func (m *LobbyMatch) MatchJoin(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, dispatcher runtime.MatchDispatcher, tick int64, state interface{}, presences []runtime.Presence) interface{} {
	lobbyState, ok := state.(*LobbyMatchState)
	if !ok {
		logger.Error("state not a valid lobby state object")
	}

	for i := 0; i < len(presences); i++ {
		lobbyState.presences[presences[i].GetSessionId()] = presences[i]
	}

	return lobbyState
}

func (m *LobbyMatch) MatchLeave(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, dispatcher runtime.MatchDispatcher, tick int64, state interface{}, presences []runtime.Presence) interface{} {
	lobbyState, ok := state.(*LobbyMatchState)
	if !ok {
		logger.Error("state not a valid lobby state object")
	}

	for i := 0; i < len(presences); i++ {
		delete(lobbyState.presences, presences[i].GetSessionId())
	}

	return lobbyState
}

func (m *LobbyMatch) MatchLoop(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, dispatcher runtime.MatchDispatcher, tick int64, state interface{}, messages []runtime.MatchData) interface{} {
	lobbyState, ok := state.(*LobbyMatchState)
	if !ok {
		logger.Error("state not a valid lobby state object")
	}

	// If we have no presences in the match according to the match state, increment the empty ticks count
	if len(lobbyState.presences) == 0 {
		lobbyState.emptyTicks++
	}

	// If the match has been empty for more than 100 ticks, end the match by returning nil
	if lobbyState.emptyTicks > 100 {
		return nil
	}

	return lobbyState
}
*/