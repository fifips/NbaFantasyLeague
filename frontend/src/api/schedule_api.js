import {Match} from "./models";

export const getSchedule = async () => {
    const dateWeekBefore = new Date();
    dateWeekBefore.setDate(dateWeekBefore.getDate() - 7);
    const dateWeekAfter = new Date();
    dateWeekAfter.setDate(dateWeekAfter.getDate() + 7);

    try {
        const res = await fetch("http://localhost:8080/schedule")
        const data = await res.json()
        return data["message"]
            .map(match => new Match(match))
            .filter(match =>  dateWeekBefore <= match.date && match.date <= dateWeekAfter)
            .sort((matchA,matchB) => matchA.date - matchB.date)
    }
    catch (err) {
        console.log('getSchedule request failed: ' + err);
        return [];
    }
}
