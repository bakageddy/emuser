<script>
	export let mediaHandle = "Social Media";
</script>
<div class="contact">
	<label for="media">
		{mediaHandle}
	</label>
	<input type="text" id="media"/>

	<br/>

	<label for="contact">
		URL
	</label>
	<input type="url" id="contact"/>
</div>

<style>
.contact {
	border-width: 1px;
	border-style: solid;
	border-radius: 4px;
}
</style>
