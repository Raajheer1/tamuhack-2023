export type Flight = {
    flight_number: string;
    departure: Airport;
    arrival: Airport;
    date: string;
    departure_time: string;
    arrival_time: string;
};

export type Airport = {
    icao: string;
    code: string;
    name: string;
    city: string;
}