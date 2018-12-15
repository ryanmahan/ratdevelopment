import { AppAuthState, AppAuthAction, AUTH_SET } from "../constants";

export const setState = (state: AppAuthState): AppAuthAction => {
    return {
        type: AUTH_SET,
        data: state
    }
};
