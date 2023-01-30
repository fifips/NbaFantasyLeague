import {useContext} from "react";
import {NotificationContext} from "../../contexts/notifContext";

export const Notifications = () => {
    const { notifications, removeNotification } = useContext(NotificationContext);

    return (
        <div className="notification-list">
            {notifications.slice(0,3).map(({ id, message }) => (
                <div key={id} className={"notification"}>
                    {/*<button className="remove-notif" onClick={() => {removeNotification(id)}}/>*/}
                    <p className="notification">
                        { message }
                    </p>
                </div>
            ))}
        </div>
    );
}