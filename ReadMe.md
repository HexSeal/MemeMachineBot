<h1> Meme Machine </h1>

![discord go logo](./images/discordgologo.png)

<h3>A Discord bot for custom meme images, written in Go</h3>



<p> Host and add it to your server(need server manage permissions):

https://discord.com/api/oauth2/authorize?client_id=766006954769776690&permissions=0&scope=bot

To run it: 
<pre>go run main.go -t (your bot token here)</pre>
</p>
</br>

<p> To create a meme: </p>
<p> - Call the bot in chat:</p>
    <pre>#meme</pre>
<p> - Choose the format(for now, the seperator is a comma)</p>
    <pre>wonka,</pre>
<p> - Add Your captions:</p>
    <pre>So tell me again, How you asked Slack instead of Googling it</pre>
<p>And the bot will post the image directly into chat!</p>

![bot output](./images/example.png)

<p>Working meme formats(more will be added):</p>
<p>- Wonka</p>