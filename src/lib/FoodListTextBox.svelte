<!-- this component will handle contacting the backend to post the food list
idea is to have it in the recipe and nutrition page -->
<script lang="ts">
    import { Input, Label, Button } from "@sveltestrap/sveltestrap";
    import type { NaturalLanguageResponseObject } from "./NutritionData";
    import { formatDateToYYYYMMDD } from "../lib/util";

    // state
    let foodListString = $state("");
    let currentSelectedDate = $state(formatDateToYYYYMMDD(new Date()));

    // cant do type safety inline, have to do it seperately
    // the ?: is the syntax for optional parameters in 'Runes Mode'
    let {
        displayCalendar,
        primaryButtonText,
        secondaryButtonText = "",
        secondaryButtonFunction = () => {},
        nutritionResponse = $bindable(), // the api response returned back to the parent component, if needed
    }: {
        displayCalendar: boolean;
        primaryButtonText: string;
        secondaryButtonText?: string;
        secondaryButtonFunction?: () => void;
        nutritionResponse: NaturalLanguageResponseObject;
    } = $props();

    let post_foodList = async (saveToDb = false) => {
        const body = {
            foodListString: foodListString,
            date: currentSelectedDate,
            saveToDb: saveToDb,
        };
        try {
            await fetch("http://localhost:8080/postFoodList", {
                method: "POST",
                headers: { "Content-Type": "application/json" },
                body: JSON.stringify(body),
            })
                .then((res) => res.json())
                .then((data) => {
                    const f = JSON.parse(data);
                    Object.assign(nutritionResponse, f);
                    // forces Svelte to rerender
                    // nutritionResponse = { ...nutritionResponse };
                    // setNutritionDisplayVisible(true);
                })
                .catch((err) => {
                    // setNutritionDisplayVisible(false);
                    console.log(err);
                    throw new Error(err);
                });
        } catch (err) {
            console.log(err);
        }
    };
</script>

<div id="food-list" class="my-2">
    <!-- this text string assortment can be turned into a reusable class -->
    <Label for="FoodListString">
        <p style="font-size: small">
            Enter a query like: <b
                >1 banana, .5 cup of white rice, 1 pound ground beef</b
            > to get the nutrition information
        </p>
        <Input
            id="FoodListString"
            type="textarea"
            placeholder="List of foods, seperated by a comma"
            bind:value={foodListString}
        />
    </Label>

    <div class="d-flex justify-content-between">
        {#if displayCalendar}
            <Label for="DatePicker" class="inline">
                <Input
                    type="date"
                    bind:value={currentSelectedDate}
                    onchange={() => console.log(currentSelectedDate)}
                />
            </Label>
        {/if}
        <Button onclick={() => post_foodList()}>{primaryButtonText}</Button>
        {#if secondaryButtonText}
            <Button onclick={secondaryButtonFunction}
                >{secondaryButtonText}</Button
            >
        {/if}
    </div>
</div>
