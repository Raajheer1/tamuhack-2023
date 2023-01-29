import { defineStore } from "pinia";

import { API } from "@/utils/api";
import {Airport, Flight} from "@/types";

interface QueryState {
    query: Query | null;
    fetching: boolean;
    hasFetched: boolean;
    loading: Promise<void> | null;
}

interface Query {
    origin: Airport;
    destination: Airport;
    data: Flight[];
}


const useSearchStore = defineStore("search", {
    state: () =>
        ({
            query: null,
            fetching: false,
            hasFetched: false,
        } as QueryState),
    getters: {
        getQuery: (state) => state.query,
    },
    actions: {
        async fetchResults() {
            this.fetching = true;
            try {
                const { data } = await API.get("/v1/search/");
                this.query = data;
            } catch (e) {
                this.query = null;
            } finally {
                this.fetching = false;
                this.hasFetched = true;
            }
        },
    },
});

export default useSearchStore;