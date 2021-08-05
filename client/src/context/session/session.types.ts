export type SessionContextData = {
    token?: string;
    username?: string;
}

export type SessionContextProperties = {
    data?: SessionContextData;
    
    exit: () => any;
    updateToken?: (token: string) => any;
}