export type Flight = {
    flight_number: string;
    departure: Airport;
    arrival: Airport;
    departure_time: string;
    arrival_time: boolean;
};

export type Airport = {
    icao: string;
    code: string;
    name: string;
    city: string;
}