import { defineStore } from "pinia";

import { API } from "@/utils/api";
import {Flight} from "@/types";

interface UserState {
    user: User | null;
    fetching: boolean;
    hasFetched: boolean;
    loading: Promise<void> | null;
}

interface User {
    id: number;
    first_name: string;
    last_name: string;
    ffnumber: string;
    trips: Flight[];
    created_at: string;
    updated_at: string;
}


const useUserStore = defineStore("user", {
    state: () =>
        ({
            user: null,
            fetching: false,
            hasFetched: false,
        } as UserState),
    getters: {
        isLoggedIn: (state) => !!state.user,
        getTrips: (state) => state.user?.trips,
        getFullName: (state) => `${state.user?.first_name} ${state.user?.last_name}`,
    },
    actions: {
        async fetchUser() {
            this.fetching = true;
            try {
                const { data } = await API.get("/v1/user/");
                this.user = data;
            } catch (e) {
                this.user = null;
            } finally {
                this.fetching = false;
                this.hasFetched = true;
            }
        },
    },
});

export default useUserStore;