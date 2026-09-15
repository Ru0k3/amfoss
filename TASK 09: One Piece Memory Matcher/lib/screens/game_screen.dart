import 'package:flutter/material.dart';
import 'package:provider/provider.dart';
import '../providers/game_provider.dart';
import '../widgets/memory_card.dart';

class GameScreen extends StatelessWidget {
  const GameScreen({super.key});

  @override
  Widget build(BuildContext context) {
    return Consumer<GameProvider>(
      builder: (context, game, child) {
        return Scaffold(
          appBar: AppBar(
            title: const Text('One Piece Memory Matcher', style: TextStyle(fontWeight: FontWeight.bold)),
            centerTitle: true,
            elevation: 0,
          ),
          body: SafeArea(
            child: Column(
              children: [
                Padding(
                  padding: const EdgeInsets.symmetric(horizontal: 24.0, vertical: 8.0),
                  child: Row(
                    mainAxisAlignment: MainAxisAlignment.spaceAround,
                    children: [
                      _buildStatColumn('Best Score', '${game.bestScore}'),
                      _buildStatColumn('Time', game.formattedTime),
                      _buildStatColumn('Moves', '${game.moves}'),
                    ],
                  ),
                ),
                Expanded(
                  child: Center(
                    child: Padding(
                      padding: const EdgeInsets.all(8.0),
                      child: AspectRatio(
                        aspectRatio: 1.0,
                        child: GridView.builder(
                          physics: const NeverScrollableScrollPhysics(),
                          gridDelegate: const SliverGridDelegateWithFixedCrossAxisCount(
                            crossAxisCount: 4,
                            crossAxisSpacing: 8.0,
                            mainAxisSpacing: 8.0,
                          ),
                          itemCount: game.cards.length,
                          itemBuilder: (context, index) {
                            return MemoryCard(
                              card: game.cards[index],
                              onTap: () => game.onCardTap(index),
                            );
                          },
                        ),
                      ),
                    ),
                  ),
                ),
                if (game.isGameOver)
                  Padding(
                    padding: const EdgeInsets.symmetric(vertical: 4.0),
                    child: Column(
                      children: [
                        Text('🎉 Congratulations! 🎉', style: Theme.of(context).textTheme.titleLarge?.copyWith(color: Colors.orange, fontWeight: FontWeight.bold)),
                        Text('Your Score: ${game.calculateScore()}', style: Theme.of(context).textTheme.titleMedium),
                      ],
                    ),
                  ),
                Padding(
                  padding: const EdgeInsets.only(top: 8.0, bottom: 16.0),
                  child: ElevatedButton.icon(
                    onPressed: game.initializeGame,
                    icon: const Icon(Icons.refresh),
                    label: const Text('Restart Game'),
                    style: ElevatedButton.styleFrom(
                      backgroundColor: Colors.orange,
                      foregroundColor: Colors.white,
                    ),
                  ),
                ),
              ],
            ),
          ),
        );
      },
    );
  }

  Widget _buildStatColumn(String label, String value) {
    return Column(
      children: [
        Text(label, style: const TextStyle(fontSize: 14, color: Colors.grey)),
        Text(value, style: const TextStyle(fontSize: 20, fontWeight: FontWeight.bold)),
      ],
    );
  }
}