<template>
  <SearchResults :flights="arr" :origin="JFK" :destination="DFW" />
</template>

<script setup lang="ts">
import SearchResults from "@/views/partials/SearchResults.vue";
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
  arrival_time: "12:30PM",
  delay_percentage: 20,
  delay_time: 10,
  price: 1234,
}
const plane1: Flight = {
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
const plane2: Flight = {
  flight_number: "AA12345",
  departure: JFK,
  arrival: DFW,
  date: "Mon 31st July",
  departure_time: "9:30PM",
  arrival_time: "10:30PM",
  delay_percentage: 10,
  delay_time: 10,
  price: 1234,
}

const arr: Flight[] = [plane, plane1, plane2]
</script>

<style scoped>

</style>