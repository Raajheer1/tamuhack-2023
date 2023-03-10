import { defineStore } from "pinia";

import { API } from "@/utils/api";
import {Airport, Flight, SearchRequest} from "@/types/index";

interface QueryState {
    query: Query | null;
    flight: Flight | null
    fetching: boolean;
    hasFetched: boolean;
    loading: Promise<void> | null;
}

interface Query {
    origin: string;
    destination: string;
    data: Flight[];
}


const useSearchStore = defineStore("search", {
    state: () =>
        ({
            query: null,
            flight: null,
            fetching: false,
            hasFetched: false,
        } as QueryState),
    getters: {
        getQuery: (state) => state.query,
    },
    actions: {
        async search(request: any) {
            this.fetching = true;
            try {
                const { data } = await API.post(`${import.meta.env.VITE_API_URL}/search`, request);
                this.query = data;
                console.log(data);
            } catch (e) {
                this.query = null;
            } finally {
                this.fetching = false;
                this.hasFetched = true;
            }
        },
        setQuery(dest: string, origin: string, flights: Flight[]) {
            const obj: Query = {
                origin: origin,
                destination: dest,
                data: flights,
            }
            this.query = obj;
        },
        saveFlight(flight: Flight) {
            this.flight = flight;
        },
        clearFlight() {
            this.flight = null;
        }
    },
});

export default useSearchStore;