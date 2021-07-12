import styled from 'styled-components';

export const HeaderContainer = styled.div`
    width: 100vw;
    height: 80px;

    background-color: #ffffff;

    display: flex;
    align-items: center;
    justify-content: space-between;
`;

export const HeaderLogo = styled.div`
    width: 50%;
    height: 100%;

    display: flex;
    align-items: center;

    img {
        width: 35px;
        height: 39px;

        margin: 0 15px 0 50px;
    }
`;

export const HeaderButtons = styled.div`
    width: 30%;
    height: 100%;

    display: flex;
    align-items: center;
    justify-content: space-between;

    padding: 0 15px;

    a {
        text-decoration: none;

        display: flex;
        align-items: center;
        justify-content: center;

        width: 350px;
        height: 40px;
        border-radius: 3px;

        margin: 0 5px;

        font-weight: bold;
    }

    a.signup {
        background-color: #F48023;
        color: #ffffff;
    }

    a.signin {
        background-color: #EAEAEA;
    }
`;