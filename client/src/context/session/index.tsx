import React, {
    createContext,
    useContext,
    
    useState,
    useCallback,
    useEffect,
} from 'react';

import { usePersistedState } from '../../util/persistedState';
import { rewoke } from '../../api/session';

import {
    SessionContextProperties,
} from './session.types';

export const SessionContext = createContext<SessionContextProperties>({} as SessionContextProperties);

export const SessionProvider: React.FC = ({
    children
}) => {
    const [token, setToken] = usePersistedState<string>("choco@jwt", "");
    const [username, setUsername] = useState<string>();  

    const exit = useCallback(() => {
        setToken("")
        setUsername("")
    }, [setToken, setUsername])

    useEffect(() => {
        rewoke(token)
            .then((response) => {
                switch(response.code) {
                    case 'SUCESS' && response.data:
                        setToken(response.data!.jwt);
                        setUsername(response.data?.username);
                        break

                    default:
                        setToken("")
                        setUsername("")
                }
            });
    }, [token, setToken, setUsername]);

    return (
        <SessionContext.Provider value={{
            data: { token, username },
            updateToken: setToken,
            exit,
        }}>
            { children }
        </SessionContext.Provider>
    )
}

export const useSession = () => useContext(SessionContext);
