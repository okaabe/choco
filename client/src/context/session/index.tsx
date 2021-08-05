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

import { AuthenticationCode } from '../../api/session/session.types';
import { useToast } from 'react-toastify';

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
        rewoke(token).then((response) => {
            if (response.code === AuthenticationCode.ACCEPTED && response.data !== null) {
                setToken(response.data!.jwt);
                setUsername(response.data?.username);
            }
        }).catch((err) => {
        })
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
