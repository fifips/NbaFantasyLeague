// user_api contains functions to communicate with backend api regarding user

// registerUser sends information about user that is trying to register to the backend
import {activateApi, isLoggedInApi, loginApi, registerApi} from "../const";

// registerUser sends POST request to backend api to provide credentials of a user trying to register
export const registerUser = async (user) => {
    const res = await fetch(registerApi, {
        method: 'POST',
        credentials: "include",
        body: JSON.stringify(user)})
    return res
}

// activateUser sends POST request to backend api with account activation code to use it
export const activateUser = async (activationCode) => {
    const res = await fetch(activateApi + activationCode, {
        method: 'POST',
        credentials: "include",
    })
    const data = await res.json()

    return data.message
}

// loginUser sends POST request to the backend api in order to login given user
export const loginUser = async ({email, password, rememberMe}) => {
    return await fetch(loginApi + `?remember=${rememberMe}`, {
        method: 'POST',
        credentials: "include",
        body: JSON.stringify({email, password})})
}

export const logoutUser = async () => {
    return await fetch("http://localhost:8080/user/logout", {
        credentials: "include",
    })
}

export const isUserLoggedIn = async () => {
    return await fetch(isLoggedInApi,{
        credentials: "include"
    })
}

