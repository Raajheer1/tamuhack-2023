export type Flight = {
    flight_number: string;
    departure: Airport;
    arrival: Airport;
    date: string;
    departure_time: string;
    arrival_time: string;
    delay_percentage: number;
    delay_time: number;
    price: number;
};

export type Airport = {
    code: string;
    name: string;
    city: string;
}