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

    console.log(theme);

    return (
        <ApplicationThemeContext.Provider value={{ theme, setTheme }}>
            <ThemeProvider theme={ theme }>
                { children }
            </ThemeProvider>
        </ApplicationThemeContext.Provider>
    );
}

export const useTheme = () => useContext(ApplicationThemeContext);