import { combineReducers } from "redux";
import { AppAuthState, AppAuthAction, AUTH_SET } from "../constants";

const AuthReducer = (state: AppAuthState = {}, action: AppAuthAction) => {
    switch (action.type) {
        case AUTH_SET:
            return action.data;
        default:
            return state;
    }
};

export const AppReducer = combineReducers({
    auth: AuthReducer
});

export interface AppState {
    auth: AppAuthState
}
