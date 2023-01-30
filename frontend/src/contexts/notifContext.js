import {createContext, useState} from "react";
import "../styles/notifications.css"
import {notifTimeout} from "../const";

export const NotificationContext = createContext({
        notifications: [],
        addNotification: () => {},
});

export const NotificationProvider = ({ children }) => {
    const [notifications, setNotifications] = useState([]);

    // TODO prevent batching on notification timeout
    const addNotification = (message) => {
        const newId = Date.now() + (5 * 1000)
        setNotifications([...notifications, { newId, message }]);
        setTimeout(() => {
            setNotifications(notifications.filter(({id, _}) => id < Date.now()));
        }, notifTimeout)
    };

    return (
        <NotificationContext.Provider value={{ notifications, addNotification }}>
            {children}
        </NotificationContext.Provider>
    );
};
