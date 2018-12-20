export const AUTH_STATUS_GUEST = 1;
export const AUTH_STATUS_LOGGEDIN = 2;
export const AUTH_STATUS_ANY = 3;
export const API_URL: string = "http://10.10.10.31:8081";

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
