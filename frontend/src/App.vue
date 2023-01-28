<template>
  <router-view />
</template>
<script setup lang="ts">
import FlightInfo from "@/components/FlightInfo.vue";
import SearchAndBookFlight from "./components/SearchAndBookFlight.vue";
import {Airport, Flight} from "@/types";

const JFK: Airport = {
  code: "JFK",
  name: "John F. Kennedy International Airport",
  city: "New York",
}

const DFW: Airport = {
  code: "DFW",
  name: "Dallas/Fort Worth International Airport",
  city: "Dallas",
}

const plane: Flight = {
  flight_number: "AA123",
  departure: JFK,
  arrival: DFW,
  date: "Mon 31st July",
  departure_time: "10:30PM",
  arrival_time: "12:30PM",
  delay_percentage: 20,
  delay_time: 10,
  price: 1234,
}
</script>