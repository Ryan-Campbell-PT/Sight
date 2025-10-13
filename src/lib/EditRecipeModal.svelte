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

    let {
        recipe = undefined,
        isOpen = false,
        onCancel = () => {},
        onSave = () => {},
    }: {
        recipe: Recipe | undefined;
        isOpen: boolean;
        onCancel: () => void;
        onSave: () => void;
    } = $props();

    const toggle = () => (isOpen = !isOpen);
</script>

<Modal {isOpen}>
    <ModalHeader>Enter Information</ModalHeader>
    {#if recipe}
        <ModalBody>
            <Form>
                <FormGroup>
                    <Label for="recipeName">Recipe name</Label>
                    <Input
                        id="recipeName"
                        type="text"
                        bind:value={recipe.recipe_name}
                        placeholder="Recipe Name"
                    />
                </FormGroup>

                <FormGroup>
                    <Label for="foodString">Food string</Label>
                    <Input
                        id="foodString"
                        type="textarea"
                        bind:value={recipe.food_string}
                        placeholder="Food String"
                    />
                </FormGroup>

                <FormGroup>
                    <Label for="servingSize">Serving Size</Label>
                    <Input
                        id="servingSize"
                        type="number"
                        bind:value={recipe.serving_size}
                        placeholder="Serving size"
                        min="0"
                    />
                </FormGroup>

                <FormGroup>
                    <Label for="numberInput">Number</Label>
                    <Input
                        id="numberInput"
                        type="checkbox"
                        bind:value={recipe.active}
                    />
                </FormGroup>
            </Form>
        </ModalBody>
    {/if}

    <ModalFooter>
        <Button color="primary" on:click={onSave}>Save</Button>
        <Button color="secondary" on:click={onCancel}>Cancel</Button>
    </ModalFooter>
</Modal>
