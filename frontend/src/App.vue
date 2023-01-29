<template>
  <router-view />
<!--  <FlightInfo :flight="plane" />-->
</template>
<script setup lang="ts">
import FlightInfo from "@/components/FlightInfoCard.vue";
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
  flight_number: "AA1234",
  departure: JFK,
  arrival: DFW,
  date: "Mon 31st July",
  departure_time: "11:30PM",
  arrival_time: "1:30AM",
  delay_percentage: 2,
  delay_time: 10,
  price: 1234,
}
</script>