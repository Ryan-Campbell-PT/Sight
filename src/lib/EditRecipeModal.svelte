<script lang="ts">
    import {
        Modal,
        ModalHeader,
        ModalBody,
        ModalFooter,
        Button,
        Form,
        FormGroup,
        Label,
        Input,
    } from "@sveltestrap/sveltestrap";
    import { onMount } from "svelte";

    let {
        recipe = $bindable(undefined),
        isOpen = false,
        onCancel = () => {},
        onSave = (r: Recipe) => {},
    }: {
        recipe: Recipe | undefined;
        isOpen: boolean;
        onCancel: () => void;
        onSave: (r: Recipe) => void;
    } = $props();

    let recipeEdit: Recipe | undefined = $state({ ...recipe } as Recipe);

    const toggle = () => (isOpen = !isOpen);
</script>

<Modal {isOpen}>
    <ModalHeader>Enter Information</ModalHeader>
    {#if recipeEdit}
        <ModalBody>
            <Form>
                <FormGroup>
                    <Label for="recipeName">Recipe name</Label>
                    <Input
                        id="recipeName"
                        type="text"
                        bind:value={recipeEdit.recipe_name}
                        placeholder="Recipe Name"
                    />
                </FormGroup>

                <FormGroup>
                    <Label for="foodString">Food string</Label>
                    <Input
                        id="foodString"
                        type="textarea"
                        bind:value={recipeEdit.food_string}
                        placeholder="Food String"
                    />
                </FormGroup>

                <FormGroup>
                    <Label for="servingSize">Serving Size</Label>
                    <Input
                        id="servingSize"
                        type="number"
                        bind:value={recipeEdit.serving_size}
                        placeholder="Serving size"
                        min="0"
                    />
                </FormGroup>

                <FormGroup>
                    <Label for="numberInput">Number</Label>
                    <Input
                        id="numberInput"
                        type="checkbox"
                        bind:value={recipeEdit.active}
                    />
                </FormGroup>
            </Form>
        </ModalBody>
    {/if}

    <ModalFooter>
        <Button color="primary" on:click={() => onSave(recipeEdit)}>Save</Button
        >
        <Button color="secondary" on:click={onCancel}>Cancel</Button>
    </ModalFooter>
</Modal>
