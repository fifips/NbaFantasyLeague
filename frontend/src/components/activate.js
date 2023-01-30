import {activateUser} from "../api/user_api";
import {useLocation} from "react-router-dom";
import {useContext, useState} from "react";
import {BouncingButton} from "./common/bouncingButton";
import "../styles/activate.css"
import {NotificationContext} from "../contexts/notifContext";

export const Activate = () => {
    const { addNotification } = useContext(NotificationContext);
    const [disabled, setDisabled] = useState(false);

    const location = useLocation();
    const params = new URLSearchParams(location.search);
    const activationCode = params.get('activation_code');

    //TODO add sending notification and redirect to home
    const handleActivation = () => {
        // activateUser(activationCode).then()
        // setDisabled(true)
        addNotification("TThis is a normal This is a normal This is a normal This is a normal his is a normal This is a " +
            "normal This is a normal This is a normal notification6This is a normal This is a normal This is a normal ")
    }

    return (
        <div className="activate-page">
            <BouncingButton disabled={disabled} onClick={handleActivation}>Activate Account</BouncingButton>
        </div>
    );
}
