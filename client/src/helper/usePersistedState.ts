import React from 'react'

export type PersistedStateResponse<T> = [
    T,
    React.Dispatch<React.SetStateAction<T>>
]

export function usePersistedState<T>(key: string, initialState: T): PersistedStateResponse<T> {
    const [state, setState] = React.useState(() => {
        const storageValue = localStorage.getItem(key)

        if (storageValue) {
            return JSON.parse(storageValue)
        }

        return initialState
    })

    React.useEffect(() => {
        localStorage.setItem(key, JSON.stringify(state))
    }, [key, state])

    return [state, setState]
}
