import {Match} from "./models";

export const getSchedule = async () => {
    const dateDayBefore = new Date();
    dateDayBefore.setDate(dateDayBefore.getDate() - 1);
    const dateWeekAfter = new Date();
    dateWeekAfter.setDate(dateWeekAfter.getDate() + 7);

    try {
        const res = await fetch("http://localhost:8080/schedule")
        const data = await res.json()
        return data["message"]
            .map(match => new Match(match))
            .filter(match =>  dateDayBefore <= match.date && match.date <= dateWeekAfter)
            .sort((matchA,matchB) => matchA.date - matchB.date)
    }
    catch (err) {
        console.log('getSchedule request failed: ' + err);
        return [];
    }
}
