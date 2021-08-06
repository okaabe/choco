import { Dispatch, SetStateAction } from "react";
import { DefaultTheme } from "styled-components";

export type ApplicationThemeProviderProperties = {
    defaultTheme: DefaultTheme;
}

export type ApplicationThemeContextProperties = {
    theme: DefaultTheme;
    setTheme: Dispatch<SetStateAction<DefaultTheme>>;
}