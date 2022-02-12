<script lang="ts">
    import {getDate, getTime, onInterval} from "./common"
    import SleepTime from "./SleepTime.svelte"
    import {Collapse, Nav, Navbar, NavbarBrand, NavbarToggler, NavItem, NavLink} from "sveltestrap";

    let now: Date = new Date()

    onInterval(() => {
        now = new Date()
    }, 1000)

    let isOpen = false;

    function handleUpdate(event) {
        isOpen = event.detail.isOpen;
    }
</script>

<Navbar dark expand="md">
    <NavbarBrand href="/">Vointini</NavbarBrand>
    <NavbarToggler on:click={() => (isOpen = !isOpen)}/>
    <Collapse {isOpen} navbar expand="md" on:update={handleUpdate}>
        <Nav class="ms-auto" navbar>
            <NavItem>
                <NavLink href="main.html">Main</NavLink>
            </NavItem>
            <NavItem>
                <NavLink href="entries.html">Entries</NavLink>
            </NavItem>
            <NavItem>
                <NavLink href="tags.html">Tags</NavLink>
            </NavItem>
            <NavItem>
                <NavLink href="tasks.html">Tasks</NavLink>
            </NavItem>
            <NavItem>
                <NavLink href="tests.html">Tests</NavLink>
            </NavItem>
            <NavItem>
                <NavLink href="resolutions.html">Resolutions</NavLink>
            </NavItem>
        </Nav>
    </Collapse>
</Navbar>

<div>{getDate(now)} {getTime(now)}</div>

<SleepTime/>

<style>

</style>