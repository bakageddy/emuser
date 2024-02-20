<script>
	import Contact from "$lib/components/Contact.svelte";
	import Skill from "$lib/components/Skill.svelte";

	function sendRequest() {
		console.log("I was clicked!");
	}

	let skill_len = 1
	function skill_increase() {
		if (skill_len < 5) skill_len += 1;
		else alert("No more than 5 Entries")
	}

	function skill_decrease() {
		if (skill_len == 0) alert("Atleast One skill is needed");
		else skill_len -= 1;
	}

	let contact_len = 1
	function contact_increase() {
		if (contact_len < 3) contact_len += 1;
		else alert("No More than 3 Contacts")
	}

	function contact_decrease() {
		if (contact_len == 0) alert("Atleast One contact is needed");
		else contact_len -= 1;
	}

</script>

<form id="resume_form">
	<label for="name">Name</label>
	<input type="text" id="name"/>
		<br/><br/>

	<label for="location">Location</label>
	<input type="text" id="location"/>
		<br/><br/>

	<label for="email">Email</label>
	<input type="text" id="email"/>
		<br/><br/>

	<label for="mobile">Mobile</label>
	<input type="tel" id="mobile"/>
		<br/><br/>

	<label for="summary">Summary</label>
	<input type="text" id="summary"/>
		<br/><br/>

	<label for="summary">Designation</label>
	<input type="text" id="designation"/>
		<br/><br/>

	<button on:click={() => contact_increase()}>
		+
	</button>
	<button on:click={() => contact_decrease()}>
		-
	</button>

	<!-- Loops over `length` times -->
	{#each {length: contact_len} as _}
		<Contact/>
	{/each}

	<br/>

	<button on:click={() => skill_increase()}>
		+
	</button>
	<button on:click={() => skill_decrease()}>
		-
	</button>

	{#each {length: skill_len} as _}
		<Skill/>
	{/each}

	<button on:click={() => sendRequest()}>
		Submit
	</button>
</form>
