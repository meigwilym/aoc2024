import fs from 'fs'

const data = fs.readFileSync('input', 'utf8').split("\n")

const rules = data.filter(s => s.includes('|')).map(s => s.split('|').map(v => parseInt(v))).sort(([a, b], [c, d]) => a-c || b-d);
const updates = data.filter(s => s.includes(',')).map(s => s.split(',').map(v => parseInt(v)));

const finalScore = updates.map(update => {
// const finalScore = [[ 75, 47, 61, 53, 29 ]].map(update => {
	const rulesForThisUpdate = rules.filter(rule => {
		return update.includes(rule[0]) && update.includes(rule[1]);
	});

	let constructedUpdate = [];
	update.forEach(u => {
		const preceding = rulesForThisUpdate.filter(([a, b]) => b === u).map(([a, b]) => a)
		constructedUpdate[preceding.length] = u
	});

	if (update.toString() !== constructedUpdate.toString()) {
		return constructedUpdate;
	}
	return [];
}).filter(ar => ar.length > 0).reduce((p, c) => {
	return c[Math.floor(c.length/2)] + p;
}, 0);

console.log(finalScore)
