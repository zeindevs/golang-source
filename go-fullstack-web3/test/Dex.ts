import { expect } from "chai";
import { ethers } from "hardhat";

describe('Dex', () => {
	it('should work', async () => {
		const [owner, otherAccount] = await ethers.getSigners()

		// deploy the test ERC20 tokens so we can actually register a new pool
		const initialSupply = ethers.parseUnits("100", 8)
		const UsdcToken = await ethers.getContractFactory("UsdcToken");
		const usdcToken = await UsdcToken.deploy(initialSupply);
		const GGToken = await ethers.getContractFactory("GGToken");
		const ggToken = await GGToken.deploy(initialSupply);

		const Dex = await ethers.getContractFactory("Dex")
		const dex = await Dex.deploy()

		await dex.registerPool(usdcToken, ggToken);

		// deposite some udsc...
		const depositAmount = ethers.parseUnits("2", 8)
		await usdcToken.approve(dex, depositAmount);
		await dex.deposit("USDCGG", usdcToken, depositAmount)
		expect(await dex.poolTokenBalances(owner, usdcToken)).to.equal(depositAmount)

		// deposite some ggtoken...
		const ggTokenDepositAmount = ethers.parseUnits("1", 8)
		await ggToken.approve(dex, ggTokenDepositAmount);
		await dex.deposit("USDCGG", ggToken, ggTokenDepositAmount)
		expect(await dex.poolTokenBalances(owner, ggToken)).to.equal(ggTokenDepositAmount)

		console.log("usdc balance before swap", await usdcToken.balanceOf(owner))
		console.log("gg token balance before swap", await ggToken.balanceOf(owner))
		const amountUSDCToSwap = ethers.parseUnits("0.5", 8)
		await usdcToken.approve(dex, amountUSDCToSwap)
		await dex.swap("USDCGG", usdcToken, amountUSDCToSwap)
		console.log("usdc balance after swap", await usdcToken.balanceOf(owner))
		console.log("gg token balance after swap", await ggToken.balanceOf(owner))

		// await dex.withdraw("USDCGG", usdcToken, depositAmount)
		// expect(await dex.poolTokenBalances(owner, usdcToken)).to.equal(0)

		// expect(await usdcToken.balanceOf(dex)).to.equal(0)
	})
})