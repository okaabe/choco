import { createGlobalStyle } from "styled-components";

export const Styles = createGlobalStyle`
    @import url('https://fonts.googleapis.com/css2?family=Roboto:wght@400;700&display=swap');

    * {
        margin: 0;
        padding: 0;
        outline: 0;
        box-sizing: border-box;
    }

    body {
        background-color: #FAFAFA;
        font-family: 'Roboto', sans-serif;
    }
`;
