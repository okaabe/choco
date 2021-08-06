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
import { toast } from 'react-toastify';

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
            if (!response.data && response.code !== AuthenticationCode.ACCEPTED) {
                return toast.error("Request to update the account information wasn't accepted")
            }

            toast.success("Updated account information")
        }).catch((err) => {
            return toast.error("Unfortunately it wasn't possible to update your account information.")
        })
    }, [token, setToken, setUsername]);

    return (
        <SessionContext.Provider value={{
            data: { token, username },
            setToken: setToken,
            exit,
        }}>
            { children }
        </SessionContext.Provider>
    )
}

export const useSession = () => useContext(SessionContext);
