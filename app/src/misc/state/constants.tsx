export const AUTH_STATUS_GUEST = 1;
export const AUTH_STATUS_LOGGEDIN = 2;
export const AUTH_STATUS_ANY = 3;
export const API_URL: string = process.env.NODE_ENV == "production" ? window.location.origin : "http://" + window.location.hostname + ":8081";
export function authorizedFetch(endpoint: string, authState: AppAuthState) {
    return fetch(API_URL + endpoint, {headers:{Authorization: "BEARER "+authState.access_token}})
}

export const AUTH_SET: string = "AUTH_SET";

export interface AppAuthState {
    authenticated?: boolean,
    userId?: number,
    userName?: string,
    access_token?: string
}

export interface AppAuthAction {
    type: string,
    data: AppAuthState
}
