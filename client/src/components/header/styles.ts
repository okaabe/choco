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
        width: 25px;
        height: 29px;

        margin: 0 15px 0 50px;
    }
`;

export const HeaderButtons = styled.div`
    width: 14%;
    height: 100%;

    display: flex;
    align-items: center;
    justify-content: space-around;

    padding: 0 15px;

    a {
        text-decoration: none;

        display: flex;
        align-items: center;
        justify-content: center;

        width: 113px;
        height: 38px;
        border-radius: 3px;

        font-size: 15px;
        font-weight: bold;

        transition: background .5s;
    }

    a.signup {
        background-color: #F48023;
        color: #ffffff;
    }

    a.signin {
        width: 82px;
        background-color: #EAEAEA;
    }

    a.signup:hover {
        background-color: #ff8421;
    }

    a.signin:hover {
        background-color: #EEEEEE;
    }

    @media (max-width: 1588px) {
        width: 18%;
    }

    @media (max-width: 1272px) {
        width: 21%;
    }

    @media (max-width: 1126px) {
        width: 25%;
    }
`;