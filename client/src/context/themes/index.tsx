import React, {
    createContext,
    useContext
} from 'react';

import { ThemeProvider, DefaultTheme } from 'styled-components';
import { ApplicationThemeContextProperties, ApplicationThemeProviderProperties } from './theme.types';

import { usePersistedState } from '../../util/persistedState';

export const ApplicationThemeContext = createContext<ApplicationThemeContextProperties>({} as ApplicationThemeContextProperties);

export const ApplicationThemeProvider: React.FC<ApplicationThemeProviderProperties> = ({
    defaultTheme,
    children
}) => {
    const [theme, setTheme] = usePersistedState<DefaultTheme>("choco@theme", defaultTheme);

    return (
        <ThemeProvider theme={ defaultTheme }>
            <ApplicationThemeContext.Provider value={{
                theme, setTheme,
            }}>
                { children }
            </ApplicationThemeContext.Provider>
        </ThemeProvider>
    );
}

export const useTheme = () => useContext(ApplicationThemeContext);