import {
    useState,
    useEffect,

    Dispatch,
    SetStateAction,
} from 'react'

export type PersistedStateResponse<T> = [T, Dispatch<SetStateAction<T>>];

export function usePersistedState<T>(key: string, initialState: T): PersistedStateResponse<T> {
    const [state, setState] = useState(() => {
        const storageValue = localStorage.getItem(key);

        if (storageValue) {
            return JSON.parse(storageValue);
        }

        return initialState;
    });

    useEffect(() => {
        localStorage.setItem(key, JSON.stringify(state))
    }, [key, state]);

    return [state, setState];
}
