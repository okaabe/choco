import styled from 'styled-components';
import Thumbnail from '../../assets/imgs/thumb.svg'

export const SignInContainer = styled.div`
    width: 100vw;
    height: 100vh;
    overflow: hidden;
`;

export const SignInContent = styled.div`
    width: 100%;
    height: 100%;

    display: flex;
    align-items: center;

    @media (max-width: 1270px) {
        justify-content: center;
    }
`;

export const SignInForm = styled.form`
    width: 35%;
    height: 100%;
    
    background-color: #fafafa;

    display: flex;
    align-items: center;
    justify-content: center;
    flex-direction: column;

    @media (max-width: 1450px) {
        width: 45%;
    }

    @media (max-width: 1270px) {
        width: 50%;
    }
`;

export const SignFormMessage = styled.div`
    width: 500px;

    h1 {
        font-size: 30px;
    }
    p {
        margin: 10px 0 10px 0;
        text-align: left;
        letter-spacing: 1px;
    }

    @media (max-width: 1471px) {
        width: 450px;
    }

    @media (max-width: 1270px) {
        width: 600px;
    }
`;

export const SignInFormInputs = styled.div`
    width: 100%;

    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;

    input {
        width: 500px;
        height: 45px;
        background-color: #ffffff;
        border-radius: 5px;
        border-style: solid;
        border-color: #EAEAEA;

        padding: 0 15px 0 15px;
        margin: 5px;
    }

    @media (max-width: 1471px) {
        input {
            width: 450px;
        }
    }

    @media (max-width: 1270px) {
        input {
            width: 600px;
        }
    }
`;

export const SignInThumbnail = styled.div`
    width: 65%;
    height: 100%;
    overflow: hidden;

    display: flex;
    align-items: flex-end;
    flex-direction: row-reverse;

    background-image: url(${Thumbnail});
    background-repeat: no-repeat;
    background-size: cover;

    @media (max-width: 1270px) {
        display: none;
    }
`;

export const SignInButton = styled.button`
    margin: 10px;

    background-color: #F48023;
    opacity: .5;

    width: 500px;
    height: 47px;

    border: none;

    transition: opacity .5s;
    cursor: pointer;

    &:hover {
        opacity: 1;
    }

    @media (max-width: 1471px) {
        width: 450px;
    }

    @media (max-width: 1270px) {
        width: 600px;
    }
`;

export const SignInErrorMessage = styled.div`
    color: #FF0000;
    margin: 5px 0 5px 0;
`