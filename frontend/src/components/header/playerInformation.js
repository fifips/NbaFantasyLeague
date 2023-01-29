import '../../styles/header/playerInformation.css';
import '../../styles/miscellaneous.css';
import {useContext} from "react";
import {AuthContext} from "../../contexts/authContext";


const LoggedInProfile = ({logout}) => {
    return (
        <div className="logged-in-profile">
            <button className="misc" onClick={logout}>
                {"Logout"}
            </button>
        </div>
    );
}

const LoggedOutProfile = ({login}) => {
    return (
        <div className="logged-out-profile">
            <button className="misc" onClick={login}>
                {"Login"}
            </button>
            <button className="misc" onClick={login}>
                {"Register"}
            </button>
        </div>
    );
}

export const UserInformation = () => {
    const {isLoggedIn, login, logout} = useContext(AuthContext);

    return (
        <div className="user-profile">
            {
                isLoggedIn
                    ? <LoggedInProfile logout={logout}/>
                    : <LoggedOutProfile login={login}/>
            }
        </div>
    );
}

