// Const file contains constant values for this program

// Backend Urls
export const backendBaseUrl = "http://localhost:8080/";
export const registerApi = backendBaseUrl + "register";
export const loginApi = backendBaseUrl + "login";
export const userBaseApi = backendBaseUrl + "user/";
export const isLoggedInApi = userBaseApi + "logged_in";
export const activateApi = userBaseApi + "activate?activation_code=";


// Notification timeout
export const notifTimeout = 5000;