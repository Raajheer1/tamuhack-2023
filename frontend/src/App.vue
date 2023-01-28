<template>
  <router-view />
  <div class="m-1">
    <FlightInfo :flight="plane" />
  </div>
</template>
<script setup lang="ts">
import FlightInfo from "@/components/FlightInfo.vue";
import {Airport, Flight} from "@/types";

const JFK: Airport = {
  icao: "KJFK",
  code: "JFK",
  name: "John F. Kennedy International Airport",
  city: "New York",
}

const DFW: Airport = {
  icao: "KDFW",
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
  arrival_time: "12:30PM"
}
</script>