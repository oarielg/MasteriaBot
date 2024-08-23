package migrations

import (
	"github.com/oarielg/MasteriaBot/internal/database"
	"github.com/oarielg/MasteriaBot/internal/models"
)

var traits []models.Trait = []models.Trait{
	{
		Name:        "Ambidextrous",
		Description: "The character is ambidextrous, thus not having disadvantage on Attack Rolls while using her offhand.",
	},
	{
		Name:        "Armor Training",
		Description: "The character trained intensely in the use of heavy armors. You gain +1 bonus to your DR while wearing Medium or Heavy armors.",
	},
	{
		Name:        "Armored Agility",
		Description: "The character prefers to fight wearing light armor to have more freedom of movement. While wearing Light armor, you gain +2 bonus to your Initiative and have advantage on all Athletics Rolls related to dodging things or acrobatics in general.",
	},
	{
		Name:        "Charming",
		Description: "The character is very good looking and endowed with great charm. You have advantage on all Manipulation Rolls.",
	},
	{
		Name:        "Blind-Fight",
		Description: "The character has trained to be able to fight properly while she cannot see. While blind, you don’t have inferiority on Attack Rolls nor you are considered helpless.",
	},
	{
		Name:        "Brawny",
		Description: "The character has a very strong and muscular body. You have +5 bonus to your Load and advantage on Athletics Rolls related to jumping, climbing, lifting heavy weights, etc., (basically any dice check related to feats of strength). ",
	},
	{
		Name:        "Charlatan",
		Description: "The character is an specialist at deceiveing others. Your Disguise power is upgraded: it now lasts for 2d hours, and its effects are only cancelled when you wish for. Additionally, whenever you use Charm or Illusion and a target makes a Composure/Insight Roll to avoid or resist your power’s effect, you can apply Effort. If you do so, the target will have inferiority on their dice check against your power.",
	},
	{
		Name:        "Coolheaded",
		Description: "The character has great self control, being able to stay calm in most situations. You have advantage on all Composure Rolls.",
	},
	{
		Name:        "Defender",
		Description: "The character’s defenses are unmatched. Your Barrier power absorbs 3 more points of damage than normal (total 9) and your Reinforce power grants +1 bonus DR (total +3).",
	},
	{
		Name:        "Destructive",
		Description: "The character is endowed with a great, destructive power. All your damaging powers (Explosion, Imbuement, Power Bolt and Power Circle) deal +1 bonus damage.",
	},
	{
		Name:        "Determined",
		Description: "The character is endowed with great will. You gain +5 bonus to your Willpower. This Trait can be taken multiple times.",
	},
	{
		Name:        "Favored Enemy",
		Description: "The character has a lot of experience hunting and fighting certain types of creatures. Choose two types of creatures (animals, celestials, demons, dragons, undead, etc.). You gain +1 bonus damage and have advantage on Insight and Perception Rolls against these types of creatures.",
	},
	{
		Name:        "Healer",
		Description: "The character possesses great healing abilities. Your Heal and Regenerate powers heal +1 point of Vitality as a bonus.",
	},
	{
		Name:        "Martial Arts",
		Description: "The character had an extensive train over martial arts. While fighting unarmed, you gain +1 bonus damage and whenever you try to parry an attack, you have advantage on the dice check.",
	},
	{
		Name:        "Menacing",
		Description: "The character has an imposing presence. You have advantage on all Intimidation Rolls.",
	},
	{
		Name:        "Mentalist",
		Description: "The character has great psychic powers. You can now communicate with your allies through your Telepathy power over a very long range (around 100 meters). Additionally, while your Telekinesis power is active, you can use it to empower your attacks (either by making your unarmed and melee attacks heavier or by making your ranged attacks projectiles travel faster), doing +1 damage bonus for each attack for the duration of the power.",
	},
	{
		Name:        "Naturalist",
		Description: "The character is intimately connected with nature and its creatures. You can now use the Polymorph and Summon Creatures powers spending 2 less Willpower, but only with animals or plants. Additionally, you can now use the Charm power on animals.",
	},
	{
		Name:        "Necromancer",
		Description: "The character has a strong connection with the dead. You can now use the powers Heal and Regenerate on undead (you can still use them on yourself). Additionally, you can now use the Summon Creatures power spending 1 less Willpower, but only with undead. Lastly, spending 1 extra Willpower (total 4) while using the Drain Life power, you can steal 1 extra point of Vitality from the target (total 3).",
	},
	{
		Name:        "Perspicacious",
		Description: "The character is very observant, insightful and analytical, possessing great perspicacy. You have advantage on all Insight Rolls.",
	},
	{
		Name:        "Powers",
		Description: "The character has special powers. Choose two powers and describe the origin of the character's powers, or how she uses them. This Trait can be taken multiple times, and every time you take it, you can choose two more powers for the character.",
	},
	{
		Name:        "Sharp Senses",
		Description: "The character’s senses are incredibly sharp. You have advantage on all Perception Rolls and gain +2 bonus to your Initiative.",
	},
	{
		Name:        "Sharpshooter",
		Description: "The character is very skilled in the use of Ranged weapons. While attacking with ranged weapons you gain +1 bonus damage and when applying Effort on the Attack Roll, your next attack will ignore all the target’s DR.",
	},
	{
		Name:        "Shield Training",
		Description: "After intense training, the character has become skilled at using her shield to deflect attacks. While wearing a shield you gain +1 bonus to your DR and whenever you try to parry an attack, you have advantage on the dice check.",
	},
	{
		Name:        "Sneak Attack",
		Description: "The character is a specialist on hitting her opponents subtly when they are distracted. While attacking with Light Melee weapons, you gain +1 bonus damage against targets that are helpless or can’t see you (attacking behind their backs also work).",
	},
	{
		Name:        "Sneaky",
		Description: "The character is very skilled at hiding her own presence. You have advantage on all Stealth Rolls.",
	},
	{
		Name:        "Swift",
		Description: "The character is endowed with great speed. You have advantage on all Escape Rolls and gain +2 bonus to your Initiative.",
	},
	{
		Name:        "Unarmed Combat",
		Description: "The character trained extensively to shape her own body as a weapon, being very skilled at unarmed combat. While fighting unarmed, you deal Damage 0 (instead of -3). Additionally, when applying Effort attacking unarmed, your next attack will ignore all the target’s DR.",
	},
	{
		Name:        "Vigorous",
		Description: "The character is extremely vigorous, capable of withstanding a huge amount of damage. You gain +5 bonus to your Vitality. This Trait can be taken multiple times.",
	},
	{
		Name:        "Weapon Master",
		Description: "The character is a master in the use of Melee weapons. While attacking with melee weapons you gain +1 bonus damage and when applying Effort on the Attack Roll, your next attack will ignore all the target’s DR.",
	},
}

func MigrateTraits() {
	_ = database.DB.Create(&traits)
}
