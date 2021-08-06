export type SessionContextData = {
    token?: string;
    username?: string;
}

export type SessionContextProperties = {
    data?: SessionContextData;
    
    exit: () => any;
    setToken: (token: string) => any;
}