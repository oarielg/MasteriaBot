package migrations

import (
	"github.com/oarielg/MasteriaBot/internal/database"
	"github.com/oarielg/MasteriaBot/internal/models"
)

var powers []models.Power = []models.Power{
	{
		Name:        "Adhesion",
		Description: "For the duration of the power, the character is able to move freely on vertical surfaces and even on the ceiling without falling. ",
		Willpower:   "2",
		Duration:    "2d Rounds",
	},
	{
		Name:        "Barrier",
		Description: "The character creates a protection barrier around the target. The barrier absorbs all damage the target takes, and will be destroyed after absorbing 6 damage or more in total.",
		Willpower:   "3",
		Duration:    "2d Rounds",
	},
	{
		Name:        "Cancel Power",
		Description: "The character can cancel an ongoing power (can’t be an instantaneous power) that was not used by her. To do that, she has to spend the amount of Willpower used to activate the power and be successful on a dice check with DL Challenging (Willpower points are still consumed on a failed check). As long as she has Willpower left, she can try multiple times.",
		Willpower:   "Variable",
		Duration:    "Instantaneous",
	},
	{
		Name:        "Charm",
		Description: "This power can be used on any humanoid within close range of the character. The target must succeed on a Composure Roll with DL Challenging - if it fails, the target becomes the character's (which will have advantage on all Social Rolls) friend for the duration of the power. Any offensive actions by the character or her allies against the target immediately breaks this effect.",
		Willpower:   "3",
		Duration:    "1d Hours",
	},
	{
		Name:        "Darkness",
		Description: "The character creates a large area of darkness. All creatures within the area (including allies, except the character) are blinded for as long as they remain within the area. Any type of lighting such as torches, candles, etc., does not work within the darkness, including a Light power.",
		Willpower:   "5",
		Duration:    "2d Rounds",
	},
	{
		Name:        "Debilitate",
		Description: "The character grants a target within close range a debiliting effect. You have to choose between Confused, Exhausted and Insecure when using the power. The effect lasts for the duration of the power, and the target can make a Composure Roll with DL Challenging to avoid being affected.",
		Willpower:   "2",
		Duration:    "2d Rounds",
	},
	{
		Name:        "Detection",
		Description: "When used, this power acts like a radar, allowing the charater to feel the presence or at least traces of what she is looking for. Examples are undead, angels, demons, magic, psionics, spirits, etc. The character isn’t capable of determining the exact location of what she is trying to detect, she can only detect the presence or not.",
		Willpower:   "2",
		Duration:    "Instantaneous",
	},
	{
		Name:        "Disguise",
		Description: "The character is able to disguise herself, changing her appearance completely for the duration of the power. The change is illusory, affecting only the character and whatever she’s carrying. Interacting with the character in any forceful way (like attacking) cancels this effect.",
		Willpower:   "2",
		Duration:    "1d Hours",
	},
	{
		Name:        "Drain Life",
		Description: "With this power, the character is able to steal the vital forces of a target within close range. Immediatelly when this power is used, and then again once per round for the duration of the power, you “steal” 2 points of Vitality from the target, that will be added to your own Vitality, without exceeding its maximum. Every round after having its Vitality stolen, the target can make a dice check with DL Demanding, ending this effect on a success.",
		Willpower:   "3",
		Duration:    "5 Rounds",
	},
	{
		Name:        "Explosion",
		Description: "The character creates a big explosion of energy (or fire, lightning, cold, etc.) in a large area hitting up to 5 targets, dealing Damage 1 to each one of them. Targets can make an Athletics Roll with DL Challenging to take only half damage. ",
		Willpower:   "5",
		Duration:    "Instantaneous",
	},
	{
		Name:        "Fear",
		Description: "The character inspires fear to a close target, which has to immediately make a Composure Roll with DL Challenging. If it fails, the target has disadvantage on all dice checks against the character for the duration of this power.",
		Willpower:   "3",
		Duration:    "2d Rounds",
	},
	{
		Name:        "Flight",
		Description: "The character is able to glide, levitate, and fly in any direction for the duration of the power. Describe how the character is capable of doing that (she grows wings, telekinesis, magnetic powers, magic, etc.).",
		Willpower:   "3",
		Duration:    "1d Hours",
	},
	{
		Name:        "Haste",
		Description: "With this power, the character enhances a close target’s speed. For the duration of this power, the target has advantage on all Physical Rolls (Athletics, Attack, Escape).",
		Willpower:   "3",
		Duration:    "2d Rounds",
	},
	{
		Name:        "Heal",
		Description: "The character touches a target (or herself), healing it for 6 Vitality points. This power does not work on undead creatures.",
		Willpower:   "3",
		Duration:    "Instantaneous",
	},
	{
		Name:        "Illusion",
		Description: "This power allows the character to create false images or sounds of things the character is familiar with. Attempting to create illusions of things the character is unfamiliar with creates imperfect illusions, allowing any creature to make an Insight Roll (DL according to the Narrator’s will) to notice the illusion and not be deceived. Any illusory image lacks solidity, allowing any creature or object to pass through it.",
		Willpower:   "3",
		Duration:    "1d Hours",
	},
	{
		Name:        "Imbuement",
		Description: "The character touches an item (weapon, armor, shield, etc.), imbuing it with power. Any creature that touches the object while this power is active takes 2 points of damage. Weapons get +2 bonus damage when you attack with them, and armors and shields inflict 2 points of damage to melee attackers each time they attack you.",
		Willpower:   "3",
		Duration:    "2d Rounds",
	},
	{
		Name:        "Inspiration",
		Description: "The character touches a target, imbuing them with great enthusiasm and bravery. For the duration of the power, the target is immune to Fear and has advantage on dice checks against Charm and Debilitate.",
		Willpower:   "3",
		Duration:    "2d Rounds",
	},
	{
		Name:        "Invisibility",
		Description: "Through this power, the character can make herself or others invisible for the duration of the power. Attacking or using a power immediately removes this effect. During a fight, just consider targets to be blind against an invisible character.",
		Willpower:   "2",
		Duration:    "2d Rounds",
	},
	{
		Name:        "Light",
		Description: "The character can make an item she touches glow bright as a torch, iluminating a medium sized area around it.",
		Willpower:   "1",
		Duration:    "2d Hours",
	},
	{
		Name:        "Natural Weapons",
		Description: "With this power the character is able to grow claws, fangs, horns, etc, and attack with them. Attacking with a Natural Weapon is considered attacking unarmed, but you deal Damage 2 with it.",
		Willpower:   "2",
		Duration:    "2d Rounds",
	},
	{
		Name:        "Polymorph",
		Description: "With this power the character is able to transform into any type of creature (another person, animal, monster, etc.), literally taking its form. When transforming into the new form, objects the character is carrying are absorbed with her, unless the new form is capable of carrying or wearing them (your choice). While in the new form, the character loses some Traits (mostly her physical ones, like Brawny) but keep others (mostly mental ones, like Perspicacy), while also receiving Traits the new form has (like being able to fly). The amount of Willpower required to use this power depends on the type of creature you want to transform into: it costs 3 Willpower for a Weak creature; 6 for a Strong one; 9 for a Leader and 12 for a Legend.",
		Willpower:   "Variable",
		Duration:    "2d Rounds",
	},
	{
		Name:        "Power Bolt",
		Description: "The character fires a bolt of energy (or fire, lightning, cold, etc.) from her hands which hits a target, dealing Damage 2. The target can make an Athletics Roll with DL Challenging to take only half damage.",
		Willpower:   "2",
		Duration:    "Instantaneous",
	},
	{
		Name:        "Power Circle",
		Description: "The character creates a large circular area around her that crackles with energy. Name one specific type of creature (like Vampires, Demons, Angels, Spirits, etc.). Creatures of the type cannot enter the area for the duration of the power, and creatures inside the area immediately take 2 points of damage when the power is used and again once per round, as long as they remain within the area. If the character leaves the area while is still active, this power is cancelled.",
		Willpower:   "3",
		Duration:    "2d Rounds",
	},
	{
		Name:        "Rage",
		Description: "With this power the character is able to enter a state of uncontrollable rage. While under this effect, you gain +1 damage bonus on melee attacks and +1 bonus to your DR. Additionally, you have advantage on all Athletics Rolls related to feats of strength.",
		Willpower:   "3",
		Duration:    "2d Rounds",
	},
	{
		Name:        "Regenerate",
		Description: "The character touches a target (or herself), healing it slowly overtime. The target recovers 2 points of Vitality immediately when the power is used, and again once per round for the duration of the power. This power does not work on undead creatures.",
		Willpower:   "3",
		Duration:    "5 Rounds",
	},
	{
		Name:        "Reinforce",
		Description: "The character reinforces the defenses of a close target, granting it +2 bonus DR for the duration of the power.",
		Willpower:   "3",
		Duration:    "2d Rounds",
	},
	{
		Name:        "Restoration",
		Description: "The character touches a target, removing a negative effect the target is being affected by. This power can remove fear, damage overtime, poison, diseases and other Conditions.",
		Willpower:   "3",
		Duration:    "Instantaneous",
	},
	{
		Name:        "Summon Creature",
		Description: "The character summons a creature to serve her. The type of creature summoned will depend on the amount of Willpower spent when activating this power: it costs 3 Willpower for a Weak creature; 6 for a Strong one; 9 for a Leader and 12 for a Legend.",
		Willpower:   "Variable",
		Duration:    "2d Rounds",
	},
	{
		Name:        "Telekinesis",
		Description: "With this power the character is able to move one mid-sized object (like a table) with her mind, accelerating or decelerating it, or leaving it suspended in the air. To levitate heavier objects you have to be successful on a dice check with DL Challenging every round, dropping the object to the ground if you fail. Additionally, a single target can be pushed using this power – doing this deals no damage, but the target is pushed 2d meters away, having to be successful on a dice check with DL Challenging to not be knocked prone.",
		Willpower:   "2",
		Duration:    "2d Rounds",
	},
	{
		Name:        "Telepathy",
		Description: "The character tries to invade the mind of a target within close range, reading its thoughts. The target can make a Composure Roll with DL Challenging to avoid having its mind read, becoming immune to this power for 24 hours on a success. On a fail, the character reads the target’s superficial thoughts (the Narrator can elaborate on them), granting the character advantage on Intimidation and Manipulation Rolls against the target.",
		Willpower:   "3",
		Duration:    "2d Rounds",
	},
	{
		Name:        "Teleport",
		Description: "This power allows the character to teleport, that is, disappear from one place and appear in another place that she has already visited. There is no limit to the distance between the two places, but they must be outdoors and on the same plane of existence. A target teleported against their will must succeed on a dice check with DL Challenging to avoid it.",
		Willpower:   "5",
		Duration:    "Instantaneous",
	},
}

func MigratePowers() {
	_ = database.DB.Create(&powers)
}
