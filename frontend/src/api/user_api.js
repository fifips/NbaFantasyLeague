// user_api contains functions to communicate with backend api regarding user

// registerUser sends information about user that is trying to register to the backend
import {activateApi, loginApi, registerApi} from "../const";

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
    //TODO add rememberMe handling (refresh token)
    const res = await fetch(loginApi, {
        method: 'POST',
        credentials: "include",
        body: JSON.stringify({email, password})})
    return res
}
//
// export const logoutUser = async () => {
//     await fetch("http://localhost:8080/user/logout", {
//         credentials: "include",
//     })
//     showNotification({ message: 'User logged out', color: 'green' });
// }
//
// export const isLoggedIn = async () => {
//     const res = await fetch("http://localhost:8080/user/logged_in",{
//         credentials: "include"
//     })
//     return res.status === 200
// }
//
